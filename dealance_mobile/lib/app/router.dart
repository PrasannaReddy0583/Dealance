import 'package:go_router/go_router.dart';
import '../features/auth/presentation/login_page.dart';

class AppRouter {
  static final router = GoRouter(
    initialLocation: '/login',
    routes: [
      GoRoute(
        path: '/login',
        builder: (_, _) => const LoginPage(),
      ),
    ],
  );
}
