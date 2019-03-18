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
    coreDirectives,
    MaterialFabComponent,
    MaterialButtonComponent,
    MaterialExpansionPanel,
    MaterialExpansionPanelSet,
    MaterialIconComponent,
    materialInputDirectives,
    NgFor,
    NgIf,
  ],
)
class ItemListComponent implements OnDestroy, OnInit {
  final ItemListService _itemListService;
  StreamSubscription _itemsSubscription;
  List<Item> itemsList = [];

  ItemListComponent(this._itemListService);

  @override
  void ngOnInit() {
    _itemListService.getAllItems();
    _itemsSubscription = _itemListService.getItemUpdateListener
        .listen((List<Item> items) => itemsList = items);
  }

  void ngOnDestroy() {
    _itemsSubscription.cancel();
  }
}
