package com.bookstore.ecommerce.app.operation.query;

import java.time.Instant;

import lombok.Getter;

public class Order {
  @Getter
  private String number;
  @Getter
  private String status;
  @Getter
  private Instant createdAt;
  @Getter
  private String customerName;
  @Getter
  private String customerPhone;
  @Getter
  private String customerDeliveryAddress;
  @Getter
  private OrderItem[] items;
}
