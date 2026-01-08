class Env {
  final String baseUrl;

  Env._(this.baseUrl);

  factory Env.production() {
    return Env._('https://api.yourcompany.dev');
  }

  factory Env.staging() {
    return Env._('https://staging-api.yourcompany.dev');
  }
}
