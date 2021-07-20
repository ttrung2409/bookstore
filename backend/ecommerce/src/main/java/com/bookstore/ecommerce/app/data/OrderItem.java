package com.bookstore.ecommerce.app.data;

import lombok.Getter;

public class OrderItem {
  @Getter private String orderId;
  @Getter private String bookId;
  @Getter private int qty;    
}