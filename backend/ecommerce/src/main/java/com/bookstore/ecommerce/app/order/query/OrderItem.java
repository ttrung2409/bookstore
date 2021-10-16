package com.bookstore.ecommerce.app.order.query;

import lombok.Data;

@Data
public class OrderItem {
  private String bookId;
  private Book book;
}
