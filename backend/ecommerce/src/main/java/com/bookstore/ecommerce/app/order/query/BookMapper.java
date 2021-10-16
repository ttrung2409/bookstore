package com.bookstore.ecommerce.app.order.query;

public class BookMapper {
  public static Book fromDataObject(com.bookstore.ecommerce.app.domain.data.Book book) {
    return Book.builder().id(book.getId()).title(book.getTitle()).subTitle(book.getSubTitle())
        .description(book.getDescription()).build();
  }
}
