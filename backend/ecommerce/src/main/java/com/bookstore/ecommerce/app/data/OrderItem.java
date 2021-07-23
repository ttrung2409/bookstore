package com.bookstore.ecommerce.app.data;

import lombok.Getter;

public class OrderItem {
  @Getter
  private String orderId;
  @Getter
  private String bookId;
  @Getter
  private String bookTitle;
  @Getter
  private String bookSubTile;
  @Getter
  private String bookDescription;
  @Getter
  private int qty;

  public OrderItem() {}

  public OrderItem(String orderId, String bookId, String bookTitle, String bookSubTilte,
      String bookDescription, int qty) {
    this.orderId = orderId;
    this.bookId = bookId;
    this.bookTitle = bookTitle;
    this.bookSubTile = bookSubTilte;
    this.bookDescription = bookDescription;
    this.qty = qty;
  }
}
