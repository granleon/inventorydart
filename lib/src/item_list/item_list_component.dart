import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_set.dart';

import '../../src/item_model.dart';

@Component(
  selector: 'app-item-list',
  styleUrls: ['item_list_component.css'],
  templateUrl: 'item_list_component.html',
  directives: [
    MaterialCheckboxComponent,
    MaterialFabComponent,
    MaterialButtonComponent,
    MaterialExpansionPanel,
    MaterialExpansionPanelSet,
    MaterialIconComponent,
    materialInputDirectives,
    NgFor,
    NgIf,
  ],
  providers: const <dynamic>[materialProviders],
)
class ItemListComponent {
  List<Item> items = [
    Item('1', 'ApoA', '444444', 'M909090', 5),
    Item('2', 'ApoB', '444445', 'M909091', 2)
  ];
}
