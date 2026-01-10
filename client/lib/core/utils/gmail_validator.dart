/// Gmail Validator Utility
/// Shared validation logic for Gmail input across the app
class GmailValidator {
  /// Validates Gmail format and domain
  /// Returns null if valid, error message if invalid
  static String? validate(String? value) {
    if (value == null || value.isEmpty) {
      return 'Gmail tidak boleh kosong';
    }

    final email = value.trim().toLowerCase();

    // Check for @ symbol
    if (!email.contains('@')) {
      return 'Format Gmail tidak valid';
    }

    // Must be @gmail.com domain
    if (!email.endsWith('@gmail.com')) {
      return 'Harus menggunakan @gmail.com';
    }

    // Gmail username validation:
    // - Starts with alphanumeric
    // - Can contain dots (but not consecutive)
    // - Min 6 chars, max 30 chars before @gmail.com
    final gmailRegex = RegExp(r'^[a-z0-9][a-z0-9\.]{4,28}[a-z0-9]@gmail\.com$');
    if (!gmailRegex.hasMatch(email)) {
      return 'Format Gmail tidak valid';
    }

    // Check for consecutive dots
    if (email.contains('..')) {
      return 'Gmail tidak boleh memiliki titik berurutan';
    }

    return null;
  }

  /// Normalizes email to lowercase and trimmed
  static String normalize(String email) {
    return email.trim().toLowerCase();
  }
}
