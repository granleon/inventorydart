import 'dart:async';

import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_auto_dismiss.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_set.dart';

@Component(
  selector: 'app-item-create',
  styleUrls: [
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
    'item_create_component.css'
  ],
  templateUrl: 'item_create_component.html',
  directives: [
    MaterialButtonComponent,
    MaterialCheckboxComponent,
    MaterialExpansionPanel,
    MaterialExpansionPanelAutoDismiss,
    MaterialExpansionPanelSet,
    MaterialFabComponent,
    MaterialIconComponent,
    materialInputDirectives,
    NgFor,
    NgIf,
  ],
  providers: const <dynamic>[materialProviders],
)
class ItemCreateComponent implements OnInit {
  @override
  Future<Null> ngOnInit() async {}
}
