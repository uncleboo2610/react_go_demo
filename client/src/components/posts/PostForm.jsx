import React from 'react';
import { Button, Form, Input, InputNumber } from 'antd';
import * as api from '../../api/index';

function PostForm() {
    const onFinish = async (values) => {
        console.log('Success:', values);
        try {
            await api.createPost(values)
        } catch (error) {
            console.log(error);
        }
      };
      
      const onFinishFailed = (errorInfo) => {
        console.log('Failed:', errorInfo);
      };
  return (
    <Form
    name="basic"
    labelCol={{ span: 8 }}
    wrapperCol={{ span: 16 }}
    style={{ maxWidth: 600 }}
    initialValues={{ remember: true }}
    onFinish={onFinish}
    onFinishFailed={onFinishFailed}
    autoComplete="off"
  >
    <Form.Item
      label="Title"
      name="Title"
      rules={[{ required: true, message: 'Please input your title!' }]}
    >
      <Input />
    </Form.Item>

    <Form.Item
      label="Body"
      name="Body"
      rules={[{ required: true, message: 'Please input your body!' }]}
    >
        <Input />
    </Form.Item>

    <Form.Item
      label="ID"
      name="ID"
      rules={[{ required: true, message: 'Please input your ID!' }]}
    >
        <InputNumber />
    </Form.Item>

    <Form.Item
      label="Name"
      name="Name"
      rules={[{ required: true, message: 'Please input your name!' }]}
    >
        <Input />
    </Form.Item>

    <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
      <Button type="primary" htmlType="submit">
        Submit
      </Button>
    </Form.Item>
  </Form>
  )
}

export default PostForm;