package com.bookstore.ecommerce.app.order.query.dto;

import lombok.Data;

@Data
public class Customer {
  private String name;
  private String phone;
  private String deliveryAddress;
}
