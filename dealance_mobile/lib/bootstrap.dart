import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'app/app.dart';
import 'core/config/env.dart';
import 'core/network/api_client.dart';

import 'features/auth/data/auth_api.dart';
import 'features/auth/data/auth_repository_impl.dart';
import 'features/auth/domain/auth_repository.dart';
import 'features/auth/presentation/bloc/auth_bloc.dart';

class Bootstrap extends StatelessWidget {
  const Bootstrap({super.key});

  @override
  Widget build(BuildContext context) {
    // 🔹 Environment
    final env = Env.production();

    // 🔹 Core infra
    final apiClient = ApiClient(env.baseUrl);

    // 🔹 Auth feature wiring
    final authApi = AuthApi(apiClient);
    final AuthRepository authRepository =
        AuthRepositoryImpl(authApi);

    return MultiRepositoryProvider(
      providers: [
        RepositoryProvider<AuthRepository>(
          create: (_) => authRepository,
        ),
      ],
      child: MultiBlocProvider(
        providers: [
          BlocProvider<AuthBloc>(
            create: (context) =>
                AuthBloc(context.read<AuthRepository>()),
          ),
        ],
        child: const App(),
      ),
    );
  }
}
