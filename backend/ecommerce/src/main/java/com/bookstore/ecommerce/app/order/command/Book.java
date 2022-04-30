package com.bookstore.ecommerce.app.order.command;

import lombok.Data;

@Data
public class Book {
  private String id;
  private String title;
  private String subtitle;
  private String description;
  private String thumbnailUrl;

  public com.bookstore.ecommerce.app.domain.data.Book toDataObject() {
    return com.bookstore.ecommerce.app.domain.data.Book.builder()
      .id(this.getId())
      .title(this.getTitle())
      .subtitle(this.getSubtitle())
      .description(this.getDescription())
      .thumbnailUrl(this.getThumbnailUrl())
      .build();
  }
}
