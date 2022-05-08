package com.bookstore.ecommerce.app.order.query;

import lombok.Builder;
import lombok.Data;

@Builder
@Data
public class Book {
  private String id;
  private String title;
  private String subtitle;
  private String description;
  private String thumbnailUrl;

  public static Book fromDataObject(com.bookstore.ecommerce.app.domain.data.Book book) {
    return Book.builder()
      .id(book.getId())
      .title(book.getTitle())
      .subtitle(book.getSubtitle())
      .description(book.getDescription())
      .thumbnailUrl(book.getThumbnailUrl())
      .build();
  }
}
