package com.bookstore.ecommerce.app.operation.command;

import lombok.Data;

@Data
public class OrderItem {
  private Book book;
  private int qty;
}
