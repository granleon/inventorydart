import 'package:angular/angular.dart';

import '../item/item_create_component.dart';
import '../item_list/item_list_component.dart';

@Component(
  selector: 'app-content',
  styleUrls: ['content_component.css'],
  templateUrl: 'content_component.html',
  directives: [ItemCreateComponent, ItemListComponent],
)
class ContentComponent {}
