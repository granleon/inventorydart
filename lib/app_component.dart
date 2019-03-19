import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

import 'src/navbar/navbar_component.dart';
import 'src/content/content_component.dart';
import 'src/item_service.dart';

@Component(
    selector: 'my-app',
    styleUrls: ['app_component.css'],
    templateUrl: 'app_component.html',
    directives: [NavbarComponent, ContentComponent],
    providers: const <dynamic>[
      materialProviders,
      ClassProvider(ItemListService)
    ])
class AppComponent {}
