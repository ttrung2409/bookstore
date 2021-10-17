package com.bookstore.ecommerce.app.order.query.model;

import lombok.Data;

@Data
public class Customer {
  private String name;
  private String phone;
  private String deliveryAddress;
}
