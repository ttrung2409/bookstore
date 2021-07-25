package com.bookstore.ecommerce.app.domain.data;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;

@AllArgsConstructor
@Builder
public class Customer {
  @Getter
  private String name;
  @Getter
  private String phone;
  @Getter
  private String deliveryAddress;
}
