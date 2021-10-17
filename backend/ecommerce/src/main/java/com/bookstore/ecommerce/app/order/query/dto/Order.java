package com.bookstore.ecommerce.app.order.query.dto;

import java.time.Instant;
import lombok.Data;

@Data
public class Order {
  private String number;
  private String status;
  private Instant createdAt;
  private Customer customer;
  private OrderItem[] items;
}
