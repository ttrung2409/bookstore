package com.bookstore.ecommerce.app.data;

import lombok.Getter;

public class Book {
  @Getter private String id;
  @Getter private String title;
  @Getter private String subTitle;
  @Getter private String description;

  public Book() {
  }

  public Book(String id, String title, String subTitle, String description) {
    this.id = id;
    this.title = title;
    this.subTitle = subTitle;
    this.description = description;
  }
}
