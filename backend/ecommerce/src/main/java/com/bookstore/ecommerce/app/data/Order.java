package com.bookstore.ecommerce.app.data;

import java.time.Instant;

import lombok.Getter;

public class Order {
  @Getter private String id;
  @Getter private String number;
  @Getter private Instant createdAt;
  @Getter private String status;
  @Getter private String customerName;
  @Getter private String customerPhone;
  @Getter private String customerDeliveryAddress;  
}
