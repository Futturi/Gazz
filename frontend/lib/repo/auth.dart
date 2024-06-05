
import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

final dio = Dio();

Future<void> saveToken(String token) async {
  final prefs = await SharedPreferences.getInstance();
  await prefs.setString('jwt_token', token);
}

Future<void> saveGroup(String group) async{
  final prefs = await SharedPreferences.getInstance();
  await prefs.setString('group', group);
}

Future<void> saveId(String id) async{
  final prefs = await SharedPreferences.getInstance();
  await prefs.setString('id', id);
}

class AuthScreenRepository{
  Future<void> Signin(String email, password) async{
    final response = await dio.post("http://localhost:8080/auth/signin",data: {'email': email, 'password': password}, options: Options(headers: {
      'Accept': 'application/json'
    }));
    if (response.statusCode == 200){
      final data =  response.data as Map<String, dynamic>;
      final token = data['token']; // получаем токен из ответа
      saveToken(token);
    }else{
      throw Exception("invalid account");
    }
  }
}

class Users{
  final String token;
  final String group;
  const Users({
    required this.token,
    required this.group
  });
}

