import React, { useState, useEffect } from 'react';
import * as api from '../../api/index';
import { Avatar, List, Button, Space } from 'antd';

function Posts () {
    const [posts, setPosts] = useState()
    const [reload,setReload] = useState(false)

    useEffect(() => { 
        const getPosts = async () => {
          try {
            const { data } = await api.fetchPosts()
            setPosts(data)
          } catch (error) {
            console.log(error);
          }
        }
        getPosts()
    }, [reload])

    const deleteHandle = async (e) => {
      const id = e.target.id;
      try {
        await api.deletePost(id)
        setReload(prev => !prev)
      } catch (error) {
        console.log(error);
      }
    }

    return (
      <div>
        <List
            itemLayout="horizontal"
            dataSource={posts}
            renderItem={(item) => (
              <List.Item>
                <List.Item.Meta
                  avatar={<Avatar src="https://joesch.moe/api/v1/random" />}
                  title={<a href="https://ant.design">{item.Title}</a>}
                  description={item.Body}
                />
                <Space>
                  <Button id={item.PostID} type="primary" >Update</Button>
                  <Button id={item.PostID} danger onClick={(e) => deleteHandle(e)}>Delete</Button>
                </Space>
              </List.Item>
            )}
          />
      </div>
    )
}

export default Posts