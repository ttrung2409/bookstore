package com.bookstore.ecommerce.app.order.command;

public class BookMapper {
  public static com.bookstore.ecommerce.app.domain.data.Book toDataObject(Book book) {
    return com.bookstore.ecommerce.app.domain.data.Book.builder().id(book.getId()).title(book.getTitle())
        .subTitle(book.getSubTitle()).description(book.getDescription()).build();
  }
}
