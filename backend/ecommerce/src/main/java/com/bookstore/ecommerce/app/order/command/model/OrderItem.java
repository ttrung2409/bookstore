package com.bookstore.ecommerce.app.order.command.model;

import lombok.Data;

@Data
public class OrderItem {
  private Book book;
  private int qty;
}
