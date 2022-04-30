package com.bookstore.ecommerce.app.order.command;

import lombok.Data;

@Data
public class Customer {
  private String name;
  private String phone;
  private String deliveryAddress;

  public com.bookstore.ecommerce.app.domain.data.Customer toDataObject() {
    return com.bookstore.ecommerce.app.domain.data.Customer.builder()
      .name(this.getName())
      .phone(this.getPhone())
      .deliveryAddress(this.getDeliveryAddress())
      .build();
  }
}
