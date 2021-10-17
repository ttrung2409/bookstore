package com.bookstore.ecommerce.app.domain.data;

import javax.persistence.Column;
import javax.persistence.Embeddable;
import javax.persistence.EmbeddedId;
import javax.persistence.Entity;
import javax.persistence.Table;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;

@Entity
@Table(name = "order_item")
@AllArgsConstructor
@Builder
public class OrderItem {
  @Getter
  @EmbeddedId
  private Key key;
  @Getter
  @Column(name = "book_title")
  private String bookTitle;
  @Getter
  @Column(name = "book_subtitle")
  private String bookSubtitle;
  @Getter
  @Column(name = "book_description")
  private String bookDescription;
  @Getter
  @Column(name = "book_thumbnail_url")
  private String bookThumbnailUrl;
  @Getter
  private int qty;

  public OrderItem() {}

  @Embeddable
  public static class Key {
    @Getter
    @Column(name = "order_id")
    private String orderId;
    @Getter
    @Column(name = "book_id")
    private String bookId;

    public Key(String orderId, String bookId) {
      this.orderId = orderId;
      this.bookId = bookId;
    }
  }
}
