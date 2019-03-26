class Item {
  String id;
  String manufacture;
  String expire;
  String lotnumber;

  Item(
    this.id,
    this.manufacture,
    this.expire,
    this.lotnumber,
  );

  factory Item.fromJson(Map<String, dynamic> item) => Item(
        item['id'],
        item['manufacture'],
        item['expire'],
        item['lotnumber'],
      );

  Map toJson() => {
        'id': id,
        'manufacture': manufacture,
        'expire': expire,
        'lotnumber': lotnumber,
      };
}
