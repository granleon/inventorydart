import 'dart:async';
import 'dart:convert';
import 'package:angular/core.dart';
import 'package:http/browser_client.dart';

import './item_model.dart';

@Injectable()
class ItemListService {
  final BrowserClient _http;
  final baseurl = 'http://localhost:9001/api/v1/item';
  static final _headers = {'Content-Type': 'application/json; charset=UTF-8'};
  Item item;
  List<Item> _itemList = [];
  final StreamController<List<Item>> _itemUpdated =
      new StreamController<List<Item>>.broadcast();

  ItemListService(this._http);

  getItem(String itemId) async {
    final url = '$baseurl/$itemId';
    final response = await _http.get(url);
    return Item.fromJson(json.decode(response.body));
  }

  getAllItems() async {
    final response = await _http.get(baseurl);
    final results = json.decode(response.body) as List;
    if (results != null) {
      _itemList = results.map((json) => Item.fromJson(json)).toList();
      _itemUpdated.add(_itemList.toList());
    } else {
      _itemUpdated.add(_itemList.toList());
    }
  }

  Stream<List<Item>> get getItemUpdateListener => _itemUpdated.stream;

  createItem(Item item) async {
    final response =
        await _http.post(baseurl, headers: _headers, body: json.encode(item));
    item = Item.fromJson(json.decode(response.body));
    _itemList.add(item);
    _itemUpdated.add(_itemList.toList());
  }

  updateItem(Item item) async {
    final updateUrl = '$baseurl/${item.id}';
    await _http.put(updateUrl, headers: _headers, body: json.encode(item));
  }

  deleteItem(String itemId) async {
    final deleteUrl = '$baseurl/$itemId';
    await _http.delete(deleteUrl, headers: _headers);
    _itemList.removeWhere((item) => item.id == itemId);
    _itemUpdated.add(_itemList.toList());
  }
}
