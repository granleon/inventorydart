class Item {
  String id;
  String lotnumber;
  String partnumber;
  String chem;
  String chemabbr;
  String expire;

  Item(this.id, this.lotnumber, this.partnumber, this.chem, this.chemabbr,
      this.expire);

  factory Item.fromJson(Map<String, dynamic> item) => Item(
      item['id'],
      item['lotnumber'],
      item['partnumber'],
      item['chem'],
      item['chemabbr'],
      item['expire']);

  Map toJson() => {
        'id': id,
        'lotnumber': lotnumber,
        'partnumber': partnumber,
        'chem': chem,
        'chemabbr': chemabbr,
        'expire': expire,
      };
}
