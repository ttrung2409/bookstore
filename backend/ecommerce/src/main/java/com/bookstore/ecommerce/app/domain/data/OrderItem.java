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
  private PartitionKey key;
  @Getter
  @Column(name = "book_title")
  private String bookTitle;
  @Getter
  @Column(name = "book_sub_title")
  private String bookSubTitle;
  @Getter
  @Column(name = "book_description")
  private String bookDescription;
  @Getter
  private int qty;

  public OrderItem() {
  }

  public static PartitionKey newPartitionKey(String orderId, String bookId) {
    return new PartitionKey(orderId, bookId);
  }

  @Embeddable
  private static class PartitionKey {
    @Getter
    @Column(name = "order_id")
    private String orderId;
    @Getter
    @Column(name = "book_id")
    private String bookId;

    public PartitionKey() {
    }

    public PartitionKey(String orderId, String bookId) {
      this.orderId = orderId;
      this.bookId = bookId;
    }
  }
}
