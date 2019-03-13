import 'package:angular/angular.dart';

import 'src/navbar/navbar_component.dart';
import 'src/content/content_component.dart';

@Component(
  selector: 'my-app',
  styleUrls: ['app_component.css'],
  templateUrl: 'app_component.html',
  directives: [NavbarComponent, ContentComponent],
)
class AppComponent {}
