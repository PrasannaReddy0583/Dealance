import '../domain/auth_repository.dart';
import '../domain/user.dart';
import 'auth_api.dart';

class AuthRepositoryImpl implements AuthRepository {
  final AuthApi api;

  AuthRepositoryImpl(this.api);

  @override
  Future<User> login(String email, String password) async {
    final data = await api.login(email, password);
    return User(
      id: data['id'],
      email: data['email'],
    );
  }
}
