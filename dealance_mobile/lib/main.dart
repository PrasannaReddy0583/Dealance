import 'package:bloc/bloc.dart';
import 'package:dealance_mobile/app/bloc_observer.dart';
import 'package:dealance_mobile/bootstrap.dart';
import 'package:flutter/material.dart';

void main() {
  Bloc.observer = AppBlocObserver();
  runApp(const Bootstrap());
}
