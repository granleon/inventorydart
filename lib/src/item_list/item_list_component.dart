import 'dart:async';

import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_set.dart';

import '../../src/item_model.dart';
import '../../src/item_service.dart';

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
class ItemListComponent implements OnDestroy {
  final ItemListService _itemListService;
  StreamSubscription _itemsSubscription;
  List<Item> itemsList = [];

  ItemListComponent(this._itemListService);

  void onActivate() {
    _itemListService.getAllItems();
    _itemsSubscription = _itemListService.getItemUpdateListener
        .listen((List<Item> items) => itemsList = items);
  }

  @override
  void ngOnDestroy() {
    _itemsSubscription.cancel();
  }
}
