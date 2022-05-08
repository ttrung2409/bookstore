package com.bookstore.ecommerce.app.domain.data;

import javax.persistence.Column;
import javax.persistence.Embeddable;
import javax.persistence.EmbeddedId;
import javax.persistence.Entity;
import javax.persistence.Table;
import org.modelmapper.ModelMapper;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;

@Entity
@Table(name = "order_item")
@Builder
@Data
public class OrderItem implements Cloneable {
  @EmbeddedId
  private Key key;
  @Column(name = "book_title")
  private String bookTitle;
  @Column(name = "book_subtitle")
  private String bookSubtitle;
  @Column(name = "book_description")
  private String bookDescription;
  @Column(name = "book_thumbnail_url")
  private String bookThumbnailUrl;
  private int qty;

  @Embeddable
  @AllArgsConstructor
  @Data
  public static class Key {
    @Column(name = "order_id")
    private String orderId;
    @Column(name = "book_id")
    private String bookId;
  }

  public OrderItem clone() {
    return new ModelMapper().map(this, OrderItem.class);
  }
}
