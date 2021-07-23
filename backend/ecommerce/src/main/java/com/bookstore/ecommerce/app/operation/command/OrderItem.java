package com.bookstore.ecommerce.app.operation.command;

import lombok.Getter;

public class OrderItem {
  @Getter
  private String bookId;
  @Getter
  private String bookTitle;
  @Getter
  private String bookSubTitle;
  @Getter
  private String bookDescription;
}
