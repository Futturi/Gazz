import 'package:shared_preferences/shared_preferences.dart';
import 'package:dio/dio.dart';

final dio = Dio();

class SubscribeRepo{
  Future<void> Subscribe(String username) async{
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('jwt_token');
    final response = await dio.post("http://localhost:8080/api/subscribe", data: {"username": username}, options: Options(headers: {
    'Accept': 'application/json',
    "Authorization": "Bearer $token"
  }));
    if (response.statusCode == 200){
      return;
    }else{
      throw Exception("Что-то пошло не так");
    }
  }

  Future<void> UnSubscribe(String username) async{
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('jwt_token');
    final response = await dio.post("http://localhost:8080/api/unsubscribe", data: {"username": username}, options: Options(headers: {
      'Accept': 'application/json',
      "Authorization": "Bearer $token"
    }));
    if (response.statusCode == 200){
      return;
    }else{
      throw Exception("Что-то пошло не так");
    }
  }
}