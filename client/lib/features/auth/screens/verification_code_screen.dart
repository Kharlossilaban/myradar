import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:iconsax/iconsax.dart';
import '../../../core/theme/app_theme.dart';
import '../../../core/widgets/custom_button.dart';
import '../../../core/utils/otp_validator.dart';
import '../../../core/services/auth_api_service.dart';
import '../../../core/network/api_exception.dart';
import 'reset_password_screen.dart';

class VerificationCodeScreen extends StatefulWidget {
  final String gmail;

  const VerificationCodeScreen({super.key, required this.gmail});

  @override
  State<VerificationCodeScreen> createState() => _VerificationCodeScreenState();
}

class _VerificationCodeScreenState extends State<VerificationCodeScreen> {
  final _formKey = GlobalKey<FormState>();
  final _codeController = TextEditingController();
  bool _isLoading = false;

  final _authService = AuthApiService();

  // Cooldown timer for resend
  int _resendCooldown = 0;
  Timer? _cooldownTimer;

  @override
  void initState() {
    super.initState();
    // Start initial cooldown after receiving OTP
    _startCooldown();
  }

  @override
  void dispose() {
    _codeController.dispose();
    _cooldownTimer?.cancel();
    super.dispose();
  }

  void _startCooldown() {
    setState(() => _resendCooldown = 60);
    _cooldownTimer?.cancel();
    _cooldownTimer = Timer.periodic(const Duration(seconds: 1), (timer) {
      if (_resendCooldown > 0) {
        setState(() => _resendCooldown--);
      } else {
        timer.cancel();
      }
    });
  }

  void _handleVerifyCode() async {
    if (_formKey.currentState!.validate()) {
      setState(() => _isLoading = true);

      // Get code input and format as PWD-XXXXXX for backend
      final codeInput = _codeController.text.trim();
      final formattedCode = 'PWD-$codeInput';

      setState(() => _isLoading = false);

      if (mounted) {
        // Navigate to reset password screen with the code
        // The actual validation will happen when user submits new password
        Navigator.push(
          context,
          MaterialPageRoute(
            builder: (context) => ResetPasswordScreen(
              gmail: widget.gmail,
              verificationCode: formattedCode,
            ),
          ),
        );
      }
    }
  }

  void _handleResendCode() async {
    if (_resendCooldown > 0) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Tunggu $_resendCooldown detik untuk kirim ulang'),
          backgroundColor: Colors.orange,
        ),
      );
      return;
    }

    try {
      // Call backend API to resend reset code
      await _authService.forgotPassword(widget.gmail);

      _startCooldown();

      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Kode verifikasi telah dikirim ulang!'),
            backgroundColor: AppTheme.successColor,
          ),
        );
      }
    } on ApiException catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(e.message),
            backgroundColor: AppTheme.errorColor,
          ),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Gagal mengirim ulang kode. Silakan coba lagi.'),
            backgroundColor: AppTheme.errorColor,
          ),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppTheme.backgroundColor,
      appBar: AppBar(backgroundColor: Colors.transparent, elevation: 0),
      body: SafeArea(
        child: SingleChildScrollView(
          padding: const EdgeInsets.all(24),
          child: Form(
            key: _formKey,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const SizedBox(height: 20),

                // Icon
                Center(
                  child: Container(
                    width: 80,
                    height: 80,
                    decoration: BoxDecoration(
                      color: AppTheme.primaryColor.withValues(alpha: 0.1),
                      borderRadius: BorderRadius.circular(20),
                    ),
                    child: const Icon(
                      Iconsax.key,
                      color: AppTheme.primaryColor,
                      size: 40,
                    ),
                  ),
                ),

                const SizedBox(height: 32),

                // Title
                Center(
                  child: Text(
                    'Verifikasi Kode',
                    style: Theme.of(context).textTheme.headlineMedium,
                  ),
                ),

                const SizedBox(height: 8),

                Center(
                  child: Text(
                    'Masukkan kode 6 digit yang telah dikirim ke\n${widget.gmail}',
                    textAlign: TextAlign.center,
                    style: Theme.of(context).textTheme.bodyMedium,
                  ),
                ),

                const SizedBox(height: 48),

                // Code Field
                Text(
                  'Kode Verifikasi (Reset Password)',
                  style: Theme.of(context).textTheme.titleSmall?.copyWith(
                    fontWeight: FontWeight.w600,
                    color: AppTheme.textPrimary,
                  ),
                ),
                const SizedBox(height: 8),
                TextFormField(
                  controller: _codeController,
                  keyboardType: TextInputType.number,
                  textAlign: TextAlign.center,
                  maxLength: 6,
                  inputFormatters: [
                    FilteringTextInputFormatter.digitsOnly,
                    LengthLimitingTextInputFormatter(6),
                  ],
                  style: const TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                    letterSpacing: 8,
                  ),
                  decoration: InputDecoration(
                    hintText: '000000',
                    counterText: '', // Hide counter
                    hintStyle: TextStyle(
                      color: AppTheme.textLight.withValues(alpha: 0.5),
                      letterSpacing: 8,
                    ),
                    prefixIcon: const Icon(
                      Iconsax.key,
                      color: AppTheme.textLight,
                    ),
                  ),
                  validator: OtpValidator.validate,
                ),

                const SizedBox(height: 32),

                // Verify Button
                CustomButton(
                  text: 'Verifikasi',
                  onPressed: _handleVerifyCode,
                  isLoading: _isLoading,
                ),

                const SizedBox(height: 24),

                // Resend Code with cooldown
                Center(
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Text(
                        'Tidak menerima kode? ',
                        style: TextStyle(color: AppTheme.textSecondary),
                      ),
                      GestureDetector(
                        onTap: _handleResendCode,
                        child: Text(
                          _resendCooldown > 0
                              ? 'Kirim Ulang ($_resendCooldown s)'
                              : 'Kirim Ulang',
                          style: TextStyle(
                            color: _resendCooldown > 0
                                ? AppTheme.textLight
                                : AppTheme.primaryColor,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
