package com.bookstore.ecommerce.app.domain.data;

import lombok.Builder;
import lombok.Getter;

@Builder
public class Customer {
  @Getter
  private String id;
  @Getter
  private String name;
  @Getter
  private String phone;
  @Getter
  private String deliveryAddress;
}
