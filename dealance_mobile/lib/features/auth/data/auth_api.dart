import '../../../core/network/api_client.dart';

class AuthApi {
  final ApiClient client;

  AuthApi(this.client);

  Future<Map<String, dynamic>> login(
      String email, String password) async {
    final response = await client.dio.post(
      '/auth/login',
      data: {
        'email': email,
        'password': password,
      },
    );
    return response.data;
  }
}
