package com.bookstore.ecommerce.app.order.query.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;

@AllArgsConstructor
@Builder
@Data
public class Book {
  private String id;
  private String title;
  private String subtitle;
  private String description;
  private String[] authors;
  private double averageRating;
  private int ratingsCount;
  private String thumbnailUrl;
  private String previewUrl;
  private int onhandQty;
  private int reservedQty;

  public static Book fromDataObject(com.bookstore.ecommerce.app.domain.data.Book book) {
    return Book.builder()
      .id(book.getId())
      .title(book.getTitle())
      .subtitle(book.getSubtitle())
      .description(book.getDescription())
      .build();
  }
}
