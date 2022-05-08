package com.bookstore.ecommerce.app.domain.data;

import org.modelmapper.ModelMapper;
import lombok.Builder;
import lombok.Data;

@Builder
@Data
public class Customer {
  private String id;
  private String name;
  private String phone;
  private String deliveryAddress;

  public Customer clone() {
    return new ModelMapper().map(this, Customer.class);
  }
}
