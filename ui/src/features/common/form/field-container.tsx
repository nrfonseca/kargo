import { Form, FormItemProps } from 'antd';
import React from 'react';
import {
  FieldValues,
  useController,
  UseControllerProps,
  UseControllerReturn
} from 'react-hook-form';

interface Props<T extends FieldValues> extends UseControllerProps<T> {
  children: (props: UseControllerReturn<T>) => React.ReactNode;
  label: string;
  formItemOptions?: Omit<FormItemProps, 'label'>;
}

export const FieldContainer = <T extends FieldValues>({
  children,
  label,
  formItemOptions,
  ...props
}: Props<T>) => {
  const controller = useController(props);

  return (
    <Form layout='vertical' component='div'>
      <Form.Item
        {...{ ...formItemOptions, label }}
        help={controller.fieldState.error?.message}
        validateStatus={controller.fieldState.error?.message ? 'error' : ''}
      >
        {children(controller)}
      </Form.Item>
    </Form>
  );
};
