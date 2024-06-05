import 'package:shared_preferences/shared_preferences.dart';
import 'package:dio/dio.dart';

class User {
  String username;
  String email;
  String password;
  String birthday;

  User(
      {required this.username, required this.email, required this.password, required this.birthday});
  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      username: json['username'],
      email: json['email'],
      password: json['password'],
      birthday: json['birthday'],
    );
  }
}

final dio = Dio();

class UsersRepo{
  Future<List<User>> GetUsers() async{
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('jwt_token');
    final response = await dio.get("http://localhost:8080/api/users", options: Options(headers: {
      'Accept': 'application/json',
      'Authorization': 'Bearer $token',
    }));
    print(response.data);
    if (response.statusCode == 200){
      final jsonResponse = response.data as Map<String, dynamic>;
      List<User> users = (jsonResponse['users'] as List)
          .map((user) => User.fromJson(user))
          .toList();
      return users;
    }else{
      throw Exception("smth wrong");
    }
  }
}