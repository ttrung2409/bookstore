package com.bookstore.ecommerce.app.order.query.dto;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Customer {
  private String id;
  private String name;
  private String phone;
  private String deliveryAddress;
}
