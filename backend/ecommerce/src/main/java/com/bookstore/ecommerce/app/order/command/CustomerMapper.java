package com.bookstore.ecommerce.app.order.command;

public class CustomerMapper {
  public static com.bookstore.ecommerce.app.domain.data.Customer toDataObject(Customer customer) {
    return com.bookstore.ecommerce.app.domain.data.Customer.builder().name(customer.getName())
        .phone(customer.getPhone()).deliveryAddress(customer.getDeliveryAddress()).build();
  }
}
