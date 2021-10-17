package com.bookstore.ecommerce.app.order.query.model;

import lombok.Data;

@Data
public class OrderItem {
  private String bookId;
  private Book book;
  private int qty;
}