package com.bookstore.ecommerce.app.operation.command;

import lombok.Data;

@Data
public class Book {
  private String id;
  private String title;
  private String subTitle;
  private String description;
}
