package com.bookstore.ecommerce.app.order.query.dto;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class OrderItem {
  private Book book;
  private int qty;

  public static OrderItem fromDataObject(com.bookstore.ecommerce.app.domain.data.OrderItem item) {
    return OrderItem.builder()
      .book(Book.builder()
        .id(item.getKey().getBookId())
        .title(item.getBookTitle())
        .subtitle(item.getBookSubtitle())
        .description(item.getBookDescription())
        .thumbnailUrl(item.getBookThumbnailUrl()).build())
      .qty(item.getQty())
      .build();
  }
}
