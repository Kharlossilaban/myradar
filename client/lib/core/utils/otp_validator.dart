/// OTP Validator Utility
/// Shared validation logic for OTP codes across the app
class OtpValidator {
  /// Validates OTP format - must be exactly 6 numeric digits
  /// Returns null if valid, error message if invalid
  static String? validate(String? value) {
    if (value == null || value.isEmpty) {
      return 'Kode tidak boleh kosong';
    }

    final code = value.trim();

    // Must be only digits
    if (!RegExp(r'^[0-9]+$').hasMatch(code)) {
      return 'Kode harus berupa angka';
    }

    // Must be exactly 6 digits
    if (code.length != 6) {
      return 'Kode harus 6 digit';
    }

    return null;
  }

  /// Cleans OTP input - removes non-digit characters
  static String clean(String value) {
    return value.replaceAll(RegExp(r'[^0-9]'), '');
  }
}
