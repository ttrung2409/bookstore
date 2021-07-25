package com.bookstore.ecommerce.app.operation.query;

import lombok.Data;

@Data
public class OrderItem {
  private String bookId;
  private Book book;
}
