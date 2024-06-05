import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:frontend/repo/auth.dart';
import 'package:frontend/repo/regist.dart';
import 'package:frontend/repo/users.dart';
import 'package:frontend/repo/subscribe.dart';

void main() => runApp(MyApp());
class MyApp extends StatelessWidget {
  late GoRouter router = GoRouter(
      routes: [GoRoute(
          path: '/',
          pageBuilder: (context, state) => MaterialPage(child: AuthScreen()),
          routes: [GoRoute(
              path: 'signup',
              pageBuilder: (context, state) => MaterialPage(child: RegistScreen())
          ),
            GoRoute(path: 'main', pageBuilder: (context, state) => MaterialPage(child: MainScreen()))]
      ),
      ]
  );
  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      theme: ThemeData(
          scaffoldBackgroundColor: const Color.fromARGB(255, 31, 31, 31),
          appBarTheme: const AppBarTheme(
              backgroundColor: Color.fromARGB(244, 31, 31, 31),
              titleTextStyle: TextStyle(
                  color: Colors.white,
                  fontSize: 20,
                  fontWeight: FontWeight.w700
              )
          ),
          textTheme: const TextTheme(
              bodySmall: TextStyle(color: Colors.white24, fontWeight: FontWeight.w400),
              bodyLarge: TextStyle(color: Colors.white, fontWeight: FontWeight.w600, fontSize: 20),
              bodyMedium: TextStyle(color: Colors.white70, fontWeight: FontWeight.w300)
          )
      ),
      title: 'App',
      routerConfig: router,
    );
  }
}

class AuthScreen extends StatefulWidget{
  @override
  _AuthScreen createState() => _AuthScreen();
}

class _AuthScreen extends State<AuthScreen>{
  final logincontoller = TextEditingController();
  final passwordcontroller = TextEditingController();
  List<Users>? token;
  String a = "123";
  bool ob = true;
  void _togglePasswordVisibility() {
    setState(() {
      ob = !ob;
    });
  }
  @override
  Widget build(BuildContext context){
    final mquery = MediaQuery.of(context);
    return Scaffold(
      appBar: AppBar(
          title: const Text("Авторизация")
      ),
      body: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(height: mquery.size.height / 10),
            Padding(padding: EdgeInsets.symmetric(horizontal: mquery.size.width / 5), child: Container(
                padding: const EdgeInsets.fromLTRB(10, 0, 10, 0),
                child: Card(
                    shape:  RoundedRectangleBorder(borderRadius: BorderRadius.circular(10.0),), color: Colors.black12,
                    child: Column(
                      children: [
                        CachedNetworkImage(imageUrl: 'https://imgtr.ee/images/2024/05/16/439e0cc848d620903cad8c6e4367c0ef.png'),
                        Padding(padding: const EdgeInsets.symmetric(horizontal: 20), child:Container(
                          child: TextField(
                            controller: logincontoller,
                            style: Theme.of(context).textTheme.bodyMedium,
                            decoration: InputDecoration(
                              hintText: 'Логин',
                              hintStyle: Theme.of(context).textTheme.bodySmall,
                              contentPadding: const EdgeInsets.symmetric(horizontal: 30.0),
                              enabledBorder: const OutlineInputBorder(
                                borderSide: BorderSide(width: 1.0, color: Colors.white24),
                              ),
                              focusedBorder: const OutlineInputBorder(
                                borderSide: BorderSide(width: 1.5, color: Colors.white),
                              ),
                            ),
                          ),
                        ),),
                        const SizedBox(height: 30,),
                        Padding(padding: const EdgeInsets.symmetric(horizontal: 20), child: Container(
                          child: TextField(
                            controller: passwordcontroller,
                            style: Theme.of(context).textTheme.bodyMedium,
                            obscureText: ob,
                            decoration: InputDecoration(
                              hintText: 'Пароль',
                              suffixIcon: IconButton(
                                icon: Icon(ob ? Icons.visibility_off : Icons.visibility),
                                onPressed: _togglePasswordVisibility,
                              ),
                              hintStyle: Theme.of(context).textTheme.bodySmall,
                              contentPadding: const EdgeInsets.symmetric(horizontal: 30.0),
                              enabledBorder: const OutlineInputBorder(
                                borderSide: BorderSide(width: 1.0, color: Colors.white24),
                              ),
                              focusedBorder: const OutlineInputBorder(
                                borderSide: BorderSide(width: 1.5, color: Colors.white),
                              ),
                            ),
                          ),
                        ),
                        ),
                        const SizedBox(height: 50,),
                        FloatingActionButton(
                            onPressed: ()async{
                              try{
                                await AuthScreenRepository().Signin(logincontoller.text, passwordcontroller.text);
                                context.go('/main');
                              } catch(e){
                                context.go('/signup');
                              }
                            },
                            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                            child: const Icon(Icons.arrow_forward, color: Colors.white70,)
                        ),
                        const SizedBox(height: 20,)
                      ],
                    )
                )
            )
            ),
            SizedBox(height: mquery.size.height / 5),
          ],
        ),
      ),
    );
  }
}

class RegistScreen extends StatefulWidget{
  @override
  _RegistScreen createState() => _RegistScreen();
}

class _RegistScreen extends State<RegistScreen>{
  final name = TextEditingController();
  final mail = TextEditingController();
  final password = TextEditingController();
  final birth = TextEditingController();
  regist? id;
  String? err;
  @override
  Widget build(BuildContext build) {
    final mquery = MediaQuery.of(context);
    return Scaffold(
        appBar: AppBar(
          title: const Text("Регистрация",),
          leading: FloatingActionButton(
            onPressed: () => Navigator.pop(context),
            backgroundColor: Theme
                .of(context)
                .scaffoldBackgroundColor,
            child: const Icon(Icons.arrow_back, color: Colors.white70,),
          ),
        ),
        body: SingleChildScrollView(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              SizedBox(height: mquery.size.height / 10,),
              Padding(padding: EdgeInsets.symmetric(
                  horizontal: mquery.size.width / 5), child: Container(
                padding: const EdgeInsets.fromLTRB(10, 0, 10, 0),
                child: Card(
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10.0),
                  ),
                  color: Colors.black12,
                  child: Column(
                    children: [
                      CachedNetworkImage(imageUrl: 'https://imgtr.ee/images/2024/05/16/439e0cc848d620903cad8c6e4367c0ef.png', width: 350,),
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 20),
                        child: TextField(
                          controller: name,
                          style: Theme
                              .of(context)
                              .textTheme
                              .bodyMedium,
                          decoration: InputDecoration(
                            hintText: 'ФИО',
                            hintStyle: Theme
                                .of(context)
                                .textTheme
                                .bodySmall,
                            contentPadding: const EdgeInsets.symmetric(
                                horizontal: 30.0),
                            enabledBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.0, color: Colors.white24),
                            ),
                            focusedBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.5, color: Colors.white),
                            ),
                          ),
                        ),),
                      const SizedBox(height: 30),
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 20),
                        child: TextField(
                          controller: mail,
                          style: Theme
                              .of(context)
                              .textTheme
                              .bodyMedium,
                          decoration: InputDecoration(
                            hintText: 'Почта',
                            hintStyle: Theme
                                .of(context)
                                .textTheme
                                .bodySmall,
                            contentPadding: const EdgeInsets.symmetric(
                                horizontal: 30.0),
                            enabledBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.0, color: Colors.white24),
                            ),
                            focusedBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.5, color: Colors.white),
                            ),
                          ),
                        ),),
                      const SizedBox(height: 30),
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 20),
                        child: TextField(
                          controller: password,
                          style: Theme
                              .of(context)
                              .textTheme
                              .bodyMedium,
                          decoration: InputDecoration(
                            hintText: 'Пароль',
                            hintStyle: Theme
                                .of(context)
                                .textTheme
                                .bodySmall,
                            contentPadding: const EdgeInsets.symmetric(
                                horizontal: 30.0),
                            enabledBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.0, color: Colors.white24),
                            ),
                            focusedBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.5, color: Colors.white),
                            ),
                          ),
                        ),),
                      const SizedBox(height: 30),
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 20),
                        child: TextField(
                          controller: birth,
                          style: Theme
                              .of(context)
                              .textTheme
                              .bodyMedium,
                          decoration: InputDecoration(
                            hintText: 'Дата рождения(формат 2024-11-24)',
                            hintStyle: Theme
                                .of(context)
                                .textTheme
                                .bodySmall,
                            contentPadding: const EdgeInsets.symmetric(
                                horizontal: 30.0),
                            enabledBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.0, color: Colors.white24),
                            ),
                            focusedBorder: const OutlineInputBorder(
                              borderSide: BorderSide(
                                  width: 1.5, color: Colors.white),
                            ),
                          ),
                        ),),
                      const SizedBox(height: 15,),
                      Text(
                        err ?? '', style: TextStyle(color: Colors.redAccent),),
                      SizedBox(height: 15,),
                      FloatingActionButton(onPressed: () async {
                        try {
                          id = await RegistRepo().signUp(
                              mail.text, password.text, name.text, birth.text);
                          context.go('/');
                        } catch (e) {
                          err = 'Такой пользователь уже существует';
                        }
                      },
                          backgroundColor: Theme
                              .of(context)
                              .scaffoldBackgroundColor,
                          child: const Icon(
                            Icons.arrow_forward, color: Colors.white70,)),
                      const SizedBox(height: 20,)
                    ],

                  ),
                ),
              ),
              ),
              SizedBox(height: mquery.size.height / 10,),
            ],
          ),)

    );
  }
}

class MainScreen extends StatefulWidget{
  @override
  _MainScreen createState() => _MainScreen();
}

class _MainScreen extends State<MainScreen>{

  List<User>? users;

  void _showSnackBar(BuildContext context, String message) {
    final snackBar = SnackBar(
      content: Text(message),
      duration: Duration(seconds: 2),
    );
    ScaffoldMessenger.of(context).showSnackBar(snackBar);
  }

  @override
  void initState(){
    super.initState();
    UsersRepo().GetUsers().then((value){
      setState(() {
        users = value;
      });
    });
  }

  @override
  Widget build(BuildContext build){
    String a = "123";
    return Scaffold(
      appBar: AppBar(
        title: const Text("Список сотрудников"),
      ),
      body: ListView.builder(
          itemBuilder: (build, idx){
              return Container(
                padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                child: Card(
                  color: Colors.grey,
                  child:  ListTile(
                    subtitle: Text(users?[idx].birthday ?? ""),
                    title: Text(users?[idx].username ?? ""),
                    trailing: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        IconButton(onPressed: () {
                          String a = users?[idx].username ?? "";
                          try{
                            SubscribeRepo().Subscribe(a);
                            _showSnackBar(context, "Вы успешно подписались на $a");
                          }catch(e){
                            _showSnackBar(context, "Что-то пошло не так");
                          }
                        }, icon: Icon(Icons.add_alert_sharp)),
                        IconButton(onPressed: (){
                          String a = users?[idx].username ?? "";
                          try{
                            SubscribeRepo().UnSubscribe(a);
                            _showSnackBar(context, "Вы успешно отписались от $a");
                          }catch(e){
                            _showSnackBar(context, "Что-то пошло не так");
                          }
                        }, icon: Icon(Icons.highlight_remove))
                      ],
                    ),
                  ) ,
                  shape:  RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10.0),
                  ),
                ),
              );
          },
          itemCount: users?.length ?? 0,
      ),
    );
  }
}


