class Item {
  String id;
  String productname;
  String partnumber;
  String lotnumber;
  int quantity;

  Item(this.id, this.productname, this.partnumber, this.lotnumber,
      this.quantity);

  factory Item.fromJson(Map<String, dynamic> item) => Item(
      item['id'],
      item['productname'],
      item['partnumber'],
      item['lotnumber'],
      item['quantity']);

  Map toJson() => {
        'id': id,
        'productname': productname,
        'partnumber': partnumber,
        'lotnumber': lotnumber,
        'quantity': quantity
      };
}
