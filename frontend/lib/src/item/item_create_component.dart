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
    // example scans
    // 0115099590228729111803271720033110M803076
    // +H6284659711L
    // +$$3200331M803076LQ

    String barcode = form.value["barcode"];

    _itemListService.createItem(barcode);
    form.reset();
  }
}
