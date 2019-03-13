import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

@Component(
  selector: 'app-navbar',
  styleUrls: [
    'package:angular_components/app_layout/layout.scss.css',
    'navbar_component.css'
  ],
  templateUrl: 'navbar_component.html',
  directives: [
    MaterialCheckboxComponent,
    MaterialFabComponent,
    MaterialButtonComponent,
    MaterialIconComponent,
    materialInputDirectives,
    NgFor,
    NgIf,
  ],
)
class NavbarComponent {}
