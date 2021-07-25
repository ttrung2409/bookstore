package com.bookstore.ecommerce.app.operation.command;

import lombok.Data;

@Data
public class Customer {
  private String name;
  private String phone;
  private String deliveryAddress;
}
