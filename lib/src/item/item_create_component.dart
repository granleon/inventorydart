import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_set.dart';

import '../item_model.dart';
import '../item_service.dart';

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
    print(form.value);
    // (01)15099590203344(11)161123(17)181130(10)M611410 // 49
    // 18 = 10  = 10 = 11
    // *+H628A409201E* // 15
    // +$$3180331M609488F
    String barcode1d = form.value["barcode1d"];
    String barcode2d = form.value["barcode2d"];
    // String unique = barcode2d.substring(5, 18);
    // String mfgdate = barcode2d.substring(19, 29);
    String expdate = barcode2d.substring(32, 38);
    String lotnumber = barcode2d.substring(42, 49);
    String partnumber = barcode1d.substring(6);

    if (barcode1d.length != 49 || barcode2d.length != 15) {
      form.reset();
    }

    Item item = Item(null, lotnumber, partnumber, null, null, expdate);
    _itemListService.createItem(item);
    form.reset();
  }
}
