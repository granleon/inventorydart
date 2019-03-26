import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_set.dart';

import '../item_service.dart';
import '../item_model.dart';

@Component(
  selector: 'app-item-create',
  styleUrls: ['item_create_component.css'],
  templateUrl: 'item_create_component.html',
  directives: [
    AutoFocusDirective,
    formDirectives,
    MaterialButtonComponent,
    MaterialExpansionPanel,
    MaterialExpansionPanelSet,
    MaterialFabComponent,
    MaterialIconComponent,
    materialInputDirectives,
    NgForm
  ],
  providers: const <dynamic>[materialProviders],
)
class ItemCreateComponent {
  final ItemListService _itemListService;
  ItemCreateComponent(this._itemListService);

  void addBarcode(NgForm form) {
    String input = form.value["barcode"];
    if (input.length != 41) {
      form.reset();
    }
    String manufacture = input.substring(18, 24);
    String expire = input.substring(26, 32);
    String lotnumber = input.substring(34);

    Item barcode = Item(null, manufacture, expire, lotnumber);

    _itemListService.createItem(barcode);
    form.reset();
  }
}
